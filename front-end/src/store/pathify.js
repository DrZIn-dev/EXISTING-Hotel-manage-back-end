import Vuex from 'vuex'
import pathify, { Payload } from 'vuex-pathify'

// debug
window.Vuex = Vuex

// decorate store mutations
var commit = Vuex.Store.prototype.commit
Vuex.Store.prototype.commit = function(type, value, options) {
    if (type.includes('@') && !(value instanceof Payload)) {
        const [mutation, property] = type.split('@')
        const payload = new Payload(mutation, '@' + property, value)
        return commit.call(this, mutation, payload, options)
    }
    return commit.call(this, type, value, options)
}

// decorate module mutations
export function updateModules(store) {
    var modules = store._modulesNamespaceMap
    Object.keys(modules).forEach(function(key) {
        const module = modules[key]
        const commit = module.context.commit
        module.context.commit = function(type, value, options) {
            if (type.includes('@') && !(options && options.root)) {
                type = key + type
                options = Object.assign({ root: true }, options || {})
            }
            return commit.call(this, type, value, options)
        }
    })
    return store
}

// TODO: upgrade register module
// const registerModule = Vuex.Store.prototype.registerModule
// Vuex.Store.prototype.registerModule = function(path, rawModule, options) {
//   const module = registerModule.call(this, path, rawModule, options)
//   // console.log('registerModule:', module)
//   return module
// }

// options
pathify.options.mapping = 'simple'

// export
export default pathify
